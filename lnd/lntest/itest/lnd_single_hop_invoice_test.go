package itest

import (
	"bytes"
	"context"
	"encoding/hex"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkt-cash/pktd/btcutil"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
	"github.com/pkt-cash/pktd/lnd/lnrpc/routerrpc"
	"github.com/pkt-cash/pktd/lnd/lntest"
	"github.com/pkt-cash/pktd/lnd/lntest/wait"
	"github.com/pkt-cash/pktd/lnd/lntypes"
	"github.com/pkt-cash/pktd/lnd/lnwire"
	"github.com/pkt-cash/pktd/lnd/record"
)

func testSingleHopInvoice(net *lntest.NetworkHarness, t *harnessTest) {
	ctxb := context.Background()

	// Open a channel with 100k satoshis between Alice and Bob with Alice being
	// the sole funder of the channel.
	ctxt, _ := context.WithTimeout(ctxb, channelOpenTimeout)
	chanAmt := btcutil.Amount(100000)
	chanPoint := openChannelAndAssert(
		ctxt, t, net, net.Alice, net.Bob,
		lntest.OpenChannelParams{
			Amt: chanAmt,
		},
	)

	// Now that the channel is open, create an invoice for Bob which
	// expects a payment of 1000 satoshis from Alice paid via a particular
	// preimage.
	const paymentAmt = 1000
	preimage := bytes.Repeat([]byte("A"), 32)
	invoice := &lnrpc.Invoice{
		Memo:      "testing",
		RPreimage: preimage,
		Value:     paymentAmt,
	}
	invoiceResp, errr := net.Bob.AddInvoice(ctxb, invoice)
	if errr != nil {
		t.Fatalf("unable to add invoice: %v", errr)
	}

	// Wait for Alice to recognize and advertise the new channel generated
	// above.
	ctxt, _ = context.WithTimeout(ctxb, defaultTimeout)
	err := net.Alice.WaitForNetworkChannelOpen(ctxt, chanPoint)
	if err != nil {
		t.Fatalf("alice didn't advertise channel before "+
			"timeout: %v", err)
	}
	err = net.Bob.WaitForNetworkChannelOpen(ctxt, chanPoint)
	if err != nil {
		t.Fatalf("bob didn't advertise channel before "+
			"timeout: %v", err)
	}

	// With the invoice for Bob added, send a payment towards Alice paying
	// to the above generated invoice.
	resp := sendAndAssertSuccess(
		t, net.Alice,
		&routerrpc.SendPaymentRequest{
			PaymentRequest: invoiceResp.PaymentRequest,
			TimeoutSeconds: 60,
			FeeLimitMsat:   noFeeLimitMsat,
		},
	)
	if hex.EncodeToString(preimage) != string(resp.PaymentPreimage) {
		t.Fatalf("preimage mismatch: expected %v, got %v", preimage,
			resp.PaymentPreimage)
	}

	// Bob's invoice should now be found and marked as settled.
	payHash := &lnrpc.PaymentHash{
		RHash: invoiceResp.RHash,
	}
	ctxt, _ = context.WithTimeout(ctxt, defaultTimeout)
	dbInvoice, errr := net.Bob.LookupInvoice(ctxt, payHash)
	if errr != nil {
		t.Fatalf("unable to lookup invoice: %v", errr)
	}
	if !dbInvoice.Settled {
		t.Fatalf("bob's invoice should be marked as settled: %v",
			spew.Sdump(dbInvoice))
	}

	// With the payment completed all balance related stats should be
	// properly updated.
	err = wait.NoError(
		assertAmountSent(paymentAmt, net.Alice, net.Bob),
		3*time.Second,
	)
	if err != nil {
		t.Fatalf(err.String())
	}

	// Create another invoice for Bob, this time leaving off the preimage
	// to one will be randomly generated. We'll test the proper
	// encoding/decoding of the zpay32 payment requests.
	invoice = &lnrpc.Invoice{
		Memo:  "test3",
		Value: paymentAmt,
	}
	ctxt, _ = context.WithTimeout(ctxt, defaultTimeout)
	invoiceResp, errr = net.Bob.AddInvoice(ctxt, invoice)
	if errr != nil {
		t.Fatalf("unable to add invoice: %v", errr)
	}

	// Next send another payment, but this time using a zpay32 encoded
	// invoice rather than manually specifying the payment details.
	sendAndAssertSuccess(
		t, net.Alice,
		&routerrpc.SendPaymentRequest{
			PaymentRequest: invoiceResp.PaymentRequest,
			TimeoutSeconds: 60,
			FeeLimitMsat:   noFeeLimitMsat,
		},
	)

	// The second payment should also have succeeded, with the balances
	// being update accordingly.
	err = wait.NoError(
		assertAmountSent(2*paymentAmt, net.Alice, net.Bob),
		3*time.Second,
	)
	if err != nil {
		t.Fatalf(err.String())
	}

	// Next send a keysend payment.
	keySendPreimage := lntypes.Preimage{3, 4, 5, 11}
	keySendHash := keySendPreimage.Hash()

	sendAndAssertSuccess(
		t, net.Alice,
		&routerrpc.SendPaymentRequest{
			Dest:           net.Bob.PubKey[:],
			Amt:            paymentAmt,
			FinalCltvDelta: 40,
			PaymentHash:    keySendHash[:],
			DestCustomRecords: map[uint64][]byte{
				record.KeySendType: keySendPreimage[:],
			},
			TimeoutSeconds: 60,
			FeeLimitMsat:   noFeeLimitMsat,
		},
	)

	// The keysend payment should also have succeeded, with the balances
	// being update accordingly.
	err = wait.NoError(
		assertAmountSent(3*paymentAmt, net.Alice, net.Bob),
		3*time.Second,
	)
	if err != nil {
		t.Fatalf(err.String())
	}

	// Now create an invoice and specify routing hints.
	// We will test that the routing hints are encoded properly.
	hintChannel := lnwire.ShortChannelID{BlockHeight: 10}
	bobPubKey := hex.EncodeToString(net.Bob.PubKey[:])
	hints := []*lnrpc.RouteHint{
		{
			HopHints: []*lnrpc.HopHint{
				{
					NodeId:                    net.Bob.PubKey[:],
					ChanId:                    hintChannel.ToUint64(),
					FeeBaseMsat:               1,
					FeeProportionalMillionths: 1000000,
					CltvExpiryDelta:           20,
				},
			},
		},
	}

	invoice = &lnrpc.Invoice{
		Memo:       "hints",
		Value:      paymentAmt,
		RouteHints: hints,
	}

	ctxt, _ = context.WithTimeout(ctxt, defaultTimeout)
	invoiceResp, errr = net.Bob.AddInvoice(ctxt, invoice)
	if errr != nil {
		t.Fatalf("unable to add invoice: %v", errr)
	}
	payreq, errr := net.Bob.DecodePayReq(ctxt, &lnrpc.PayReqString{PayReq: invoiceResp.PaymentRequest})
	if errr != nil {
		t.Fatalf("failed to decode payment request %v", errr)
	}
	if len(payreq.RouteHints) != 1 {
		t.Fatalf("expected one routing hint")
	}
	routingHint := payreq.RouteHints[0]
	if len(routingHint.HopHints) != 1 {
		t.Fatalf("expected one hop hint")
	}
	hopHint := routingHint.HopHints[0]
	if hopHint.FeeProportionalMillionths != 1000000 {
		t.Fatalf("wrong FeeProportionalMillionths %v",
			hopHint.FeeProportionalMillionths)
	}
	if hex.EncodeToString(hopHint.NodeId) != bobPubKey {
		t.Fatalf("wrong NodeId %v",
			hopHint.NodeId)
	}
	if hopHint.ChanId != hintChannel.ToUint64() {
		t.Fatalf("wrong ChanId %v",
			hopHint.ChanId)
	}
	if hopHint.FeeBaseMsat != 1 {
		t.Fatalf("wrong FeeBaseMsat %v",
			hopHint.FeeBaseMsat)
	}
	if hopHint.CltvExpiryDelta != 20 {
		t.Fatalf("wrong CltvExpiryDelta %v",
			hopHint.CltvExpiryDelta)
	}

	ctxt, _ = context.WithTimeout(ctxb, channelCloseTimeout)
	closeChannelAndAssert(ctxt, t, net, net.Alice, chanPoint, false)
}

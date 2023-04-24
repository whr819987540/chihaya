// Package frontend provides implementations and abstractions for tracker logic,
// parsing and request handling over multiple protocols, such as UDP and HTTP.
package frontend

import (
	"context"

	"github.com/whr819987540/chihaya/bittorrent"
)

// TrackerLogic is the interface used by a frontend in order to: (1) generate a
// response from a parsed request, and (2) asynchronously observe anything
// after the response has been delivered to the client.
type TrackerLogic interface {
	// HandleAnnounce generates a response for an Announce.
	//
	// Returns the updated context, the generated AnnounceResponse and no error
	// on success; nil and error on failure.
	HandleAnnounce(context.Context, *bittorrent.AnnounceRequest) (context.Context, *bittorrent.AnnounceResponse, error)

	// AfterAnnounce does something with the results of an Announce after it
	// has been completed.
	AfterAnnounce(context.Context, *bittorrent.AnnounceRequest, *bittorrent.AnnounceResponse)

	// HandleScrape generates a response for a Scrape.
	//
	// Returns the updated context, the generated AnnounceResponse and no error
	// on success; nil and error on failure.
	HandleScrape(context.Context, *bittorrent.ScrapeRequest) (context.Context, *bittorrent.ScrapeResponse, error)

	// AfterScrape does something with the results of a Scrape after it has been completed.
	AfterScrape(context.Context, *bittorrent.ScrapeRequest, *bittorrent.ScrapeResponse)
}

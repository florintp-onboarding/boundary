// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bsr

import (
	"context"
	"time"
)

// BaseSummary contains the common fields of all summary types.
type BaseSummary struct {
	Id        string
	StartTime time.Time
	EndTime   time.Time
	Errors    error
}

type Summary interface {
	// GetId returns the Id of the summary file.
	GetId() string
	// GetStartTime returns the start time using a monotonic clock of the summary.
	GetStartTime() time.Time
	// GetEndTime returns the end time using a monotonic clock of the summary.
	GetEndTime() time.Time
}

// GetId returns the Id of the summary file.
func (b *BaseSummary) GetId() string {
	return b.Id
}

// GetStartTime returns the start time using a monotonic clock of the summary.
func (b *BaseSummary) GetStartTime() time.Time {
	return b.StartTime
}

// GetEndTime returns the end time using a monotonic clock of the summary.
func (b *BaseSummary) GetEndTime() time.Time {
	return b.EndTime
}

// BaseSessionSummary encapsulates data for a session, including its session id, connection count,
// and start/end time using a monotonic clock
type BaseSessionSummary struct {
	Id              string
	ConnectionCount uint64
	StartTime       time.Time
	EndTime         time.Time
	Errors          error
}

type SessionSummary interface {
	Summary
	// GetConnectionCount returns the connection count.
	GetConnectionCount() uint64
}

func AllocSessionSummary(_ context.Context) Summary {
	return &BaseSessionSummary{}
}

// GetId returns the Id of the summary file.
func (b *BaseSessionSummary) GetId() string {
	return b.Id
}

// GetStartTime returns the start time using a monotonic clock of the summary.
func (b *BaseSessionSummary) GetStartTime() time.Time {
	return b.StartTime
}

// GetEndTime returns the end time using a monotonic clock of the summary.
func (b *BaseSessionSummary) GetEndTime() time.Time {
	return b.EndTime
}

// GetConnectionCount returns the connection count.
func (b *BaseSessionSummary) GetConnectionCount() uint64 {
	return b.ConnectionCount
}

// BaseChannelSummary encapsulates data for a channel, including its id, channel type,
// start/end time using a monotonic clock, and the bytes up/ down seen on this channel
type BaseChannelSummary struct {
	Id                    string
	ConnectionRecordingId string
	StartTime             time.Time
	EndTime               time.Time
	BytesUp               uint64
	BytesDown             uint64
	ChannelType           string
}

type ChannelSummary interface {
	Summary
	// GetConnectionRecordingId returns the connection recording id of the channel.
	GetConnectionRecordingId() string
	// GetBytesUp returns upload bytes.
	GetBytesUp() uint64
	// BytesDown returns download bytes.
	GetBytesDown() uint64
	// GetChannelType the type of summary channel.
	GetChannelType() string
}

func AllocChannelSummary(_ context.Context) Summary {
	return &BaseChannelSummary{}
}

// GetId returns the Id of the summary file.
func (b *BaseChannelSummary) GetId() string {
	return b.Id
}

// GetId returns the Id of the summary file.
func (b *BaseChannelSummary) GetConnectionRecordingId() string {
	return b.ConnectionRecordingId
}

// GetStartTime returns the start time using a monotonic clock of the summary.
func (b *BaseChannelSummary) GetStartTime() time.Time {
	return b.StartTime
}

// GetEndTime returns the end time using a monotonic clock of the summary.
func (b *BaseChannelSummary) GetEndTime() time.Time {
	return b.EndTime
}

// GetBytesUp returns upload bytes.
func (b *BaseChannelSummary) GetBytesUp() uint64 {
	return b.BytesUp
}

// GetBytesDown returns download bytes.
func (b *BaseChannelSummary) GetBytesDown() uint64 {
	return b.BytesDown
}

// GetChannelType returns the type of summary channel.
func (b *BaseChannelSummary) GetChannelType() string {
	return b.ChannelType
}

// BaseConnectionSummary encapsulates data for a connection, including its connection id, channel count,
// start/end time using a monotonic clock, and the aggregate bytes up/ down of its channels
type BaseConnectionSummary struct {
	Id           string
	ChannelCount uint64
	StartTime    time.Time
	EndTime      time.Time
	BytesUp      uint64
	BytesDown    uint64
	Errors       error
}

type ConnectionSummary interface {
	Summary
	// GetChannelCount returns the channel count.
	GetChannelCount() uint64
	// GetBytesUp returns upload bytes.
	GetBytesUp() uint64
	// BytesDown returns download bytes.
	GetBytesDown() uint64
}

func AllocConnectionSummary(_ context.Context) Summary {
	return &BaseConnectionSummary{}
}

// GetChannelCount returns the channel count.
func (b *BaseConnectionSummary) GetChannelCount() uint64 {
	return b.ChannelCount
}

// GetId returns the Id of the summary file.
func (b *BaseConnectionSummary) GetId() string {
	return b.Id
}

// GetStartTime returns the start time using a monotonic clock of the summary.
func (b *BaseConnectionSummary) GetStartTime() time.Time {
	return b.StartTime
}

// GetEndTime returns the end time using a monotonic clock of the summary.
func (b *BaseConnectionSummary) GetEndTime() time.Time {
	return b.EndTime
}

// GetBytesUp BaseConnectionSummary upload bytes.
func (b *BaseConnectionSummary) GetBytesUp() uint64 {
	return b.BytesUp
}

// GetBytesDown returns download bytes.
func (b *BaseConnectionSummary) GetBytesDown() uint64 {
	return b.BytesDown
}

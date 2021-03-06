package ripo

import (
	"fmt"
	"log"
	"strings"
)

var errorDispatcher = func(request ExtendedRequest, rpcErr RPCError) {
	parts := []string{
		fmt.Sprintf("Code=%v", rpcErr.Code()),
		fmt.Sprintf("Message=%#v", rpcErr.Message()),
	}
	if rpcErr.Cause() != nil {
		parts = append(parts, fmt.Sprintf("Cause=%#v", rpcErr.Cause().Error()))
	}
	if len(rpcErr.Details()) > 0 {
		parts = append(parts, fmt.Sprintf("Details=%#v", rpcErr.Details()))
	}
	tbLines := []string{
		"Traceback:",
	}
	for _, record := range rpcErr.Traceback(request.HandlerName()).Records() {
		tbLines = append(tbLines, fmt.Sprintf(
			"\tFile %v, Func %v, Line %v",
			record.File(),
			record.FunctionLocal(),
			record.Line(),
		))
	}
	parts = append(parts, strings.Join(tbLines, "\n"))
	log.Printf("RPCError: %v, \n", strings.Join(parts, ", "))
}

func SetErrorDispatcher(dispatcher func(request ExtendedRequest, rpcErr RPCError)) {
	errorDispatcher = dispatcher
}

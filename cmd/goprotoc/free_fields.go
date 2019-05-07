package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/jhump/protoreflect/desc"
)

func doPrintFreeFieldNumbers(fds []*desc.FileDescriptor, w io.Writer) {
	for _, fd := range fds {
		for _, md := range fd.GetMessageTypes() {
			printMessageFreeFields(md, w)
		}
	}
}

func printMessageFreeFields(md *desc.MessageDescriptor, w io.Writer) {
	for _, nested := range md.GetNestedMessageTypes() {
		printMessageFreeFields(nested, w)
	}

	unused := computeFreeRanges(md)
	fmt.Fprintf(w, "%- 35s free:", md.GetFullyQualifiedName())
	for _, r := range unused {
		if r.end == maxTag {
			fmt.Fprintf(w, " %d-INF", r.start)
		} else if r.start == r.end {
			fmt.Fprintf(w, " %d", r.start)
		} else {
			fmt.Fprintf(w, " %d-%d", r.start, r.end)
		}
	}
	fmt.Fprintln(w)
}

func computeFreeRanges(md *desc.MessageDescriptor) []tagRange {
	var used []tagRange
	// compute all used ranges
	for _, fd := range md.GetFields() {
		used = append(used, tagRange{start: fd.GetNumber(), end: fd.GetNumber()})
	}
	for _, rr := range md.AsDescriptorProto().GetReservedRange() {
		used = append(used, tagRange{start: rr.GetStart(), end: rr.GetEnd()-1})
	}
	for _, extr := range md.GetExtensionRanges() {
		used = append(used, tagRange{start: extr.Start, end: extr.End})
	}
	// sort
	sort.Slice(used, func(i, j int) bool {
		return used[i].start < used[j].start
	})
	// now compute the inverse (unused ranges)
	unused := make([]tagRange, 0, len(used)+1)
	last := int32(0)
	for _, r := range used {
		if r.start <= last+1 {
			last = r.end
			continue
		}
		unused = append(unused, tagRange{start: last+1, end: r.start-1})
		last = r.end
	}
	if last < maxTag {
		unused = append(unused, tagRange{start: last+1, end: maxTag})
	}
	return unused
}

type tagRange struct {
	start, end int32
}
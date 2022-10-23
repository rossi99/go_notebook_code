package main

// Escape Analysis (EA) is used by compiler, to determine if a value should be constructed on the stack (default) or
// heap. The construction of a value only happens once and EA decided whether it will be - stack or heap. Only
// construction on the heap is called an allocation.

// Understanding EA is about understanding value ownership. When a value is constructed within a func, that func owns
// the value. Then ask, does the value still need to exist when the owning func returns? If no, the value can be
// constructed on the stack, otherwise, on the heap.

// Code below demonstrates this:
type user struct {
	name  string
	email string
}

// stayOnStack func is using value semantics to return a user value back to the caller, i.e , the caller gets its own
// copy of the user value being constructed.
func stayOnStack() user {
	u := user{
		name:  "John Doe",
		email: "john@doe.com",
	}
	return u
}

// When stayOnStack is called and returns, the user value it constructs no longer needs to exist as caller gets
// their own copy. No allocation.

// escapeToHeap used pointer semantics to return user value back to caller, i.e, caller gets shared access (address) to
// the user value constructed.
func escapeToHeap() *user {
	u := user{
		name:  "John Doe",
		email: "john@doe.com",
	}
	return &u
}

// When escapeToHeap is called and returns, the user value it constructs does still need to exist, since caller is
// getting shared access <- construction can't happen on stack so must be on heap. Allocation.

// If the user value here was constructed on the stack, the caller would get a copy of the stack address from frame
// below and integrity would be lost. Once control is back to calling func, memory on stack where user value exists is
// reusable. As soon as calling func makes call to other func, a new frame is sliced & memory is overridden <- destroys
// shared value.


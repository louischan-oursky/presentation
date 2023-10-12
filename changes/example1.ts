// Suppose we have this function to return a 0-based index.
function getIndex<T>(a: T[]): number {
  return 0;
}

// Some time later, we discover that we need it to return a 1-based index.
// If you ever do this, you are risking to break every consumer of this function!
function getIndex<T>(a: T[]): number {
	return 1;
}

// A better to introduce this to is rename getIndex to getIndexZeroBased,
// and add getIndexOneBased.
// Note that in this case, we cannot "Change the return type",
// as we are still returning a number.
function getIndexZeroBased<T>(a: T[]): number {
  return 0;
}

function getIndexOneBased<T>(a: T[]): number {
  return 1;
}

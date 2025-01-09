# sync.Map LoadOrStore Inconsistency in Go

This repository demonstrates a subtle inconsistency in Go's `sync.Map`'s `LoadOrStore` method.  Under certain conditions, `LoadOrStore` may return a new value without actually updating the internal map.

The issue is showcased in `bug.go`. The code attempts to use `LoadOrStore` to update a value in the `sync.Map`. Although `LoadOrStore` successfully returns the new value, the subsequent `Load` call shows that the map hasn't actually been updated.  This behavior is unexpected and can lead to hard-to-debug issues.

`bugSolution.go` provides a workaround which utilizes the traditional `Load` and `Store` methods to achieve reliable updates.

This inconsistency likely stems from internal optimization strategies within `sync.Map`, which prioritize speed over atomicity in certain edge cases. Therefore, for critical applications demanding absolute reliability, it's advised to avoid relying on `LoadOrStore` for updates and utilize a more explicit approach.
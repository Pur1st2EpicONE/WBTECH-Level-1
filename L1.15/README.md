## L1.15

This Go snippet demonstrates two approaches to extracting a substring from a larger string while controlling memory usage.

The someFunc function takes a substring of a large string using a simple slice (v[:100]). This approach is straightforward and efficient in terms of execution, but the substring shares the same underlying array as the original string, keeping the full 1KB allocation in memory even if only the first 100 bytes are needed.

The fixedFunc function explicitly copies the first 100 bytes into a new string using string([]byte(v[:100])). This ensures that the resulting string is independent of the original large string, avoiding unnecessary memory retention.

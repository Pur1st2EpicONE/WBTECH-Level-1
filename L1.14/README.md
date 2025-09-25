## L1.14

This Go snippet demonstrates two different approaches to determining the runtime type of a variable passed as an empty interface.

The assertType function uses a type switch, directly matching against concrete types like int, string, and bool. For channels, it requires explicitly listing each possible channel type, making it less flexible but very straightforward.

The assertReflection function uses the reflect package, checking the Kind of the type at runtime. This approach allows generic detection of any channel type without needing to enumerate them explicitly, while still handling basic types like integers, strings, and booleans. However, since reflect.TypeOf() returns nil, it requires an explicit nil check before inspecting the type to avoid runtime errors.

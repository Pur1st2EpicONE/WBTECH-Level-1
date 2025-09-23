## L1.10

This Go snippet demonstrates grouping temperatures into ranges using a map.

The program defines a single function, group, which iterates over an array of temperatures.

For each temperature, it calculates the start of its 10-degree range by dividing the value by 10, converting it to an integer, and multiplying by 10. The temperature is then appended to a slice in the map corresponding to that range.

The main function calls group and prints the resulting map, showing temperatures organized by ranges.
# memballast
Memory ballast for rarer GC calls

[More info...](https://medium.com/clean-code-channel/go-memory-ballast-dec0c04830b1)

## usage 
 + In main file
   ```go
   import (
      _ "github.com/vl4deee11/memballast"
   )
   ```
 + Set ENV `MEMBALLAST_BYTE`

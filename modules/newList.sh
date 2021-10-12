#!/bin/sh

TYPE='Person'

cat > ./modules/${TYPE}List.go <<EOL
package modules
type ${TYPE}List []*${TYPE}
type ${TYPE}ToBool func(*${TYPE}) bool
func (al ${TYPE}List)Filter(f ${TYPE}ToBool) ${TYPE}List {
   var ret ${TYPE}List
   for _, a := range al {
      if f(a) {
         ret = append(ret, a)
      }
   }
   return ret
}
EOL
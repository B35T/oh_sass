# oh_sass SASS and Stylus compile to css on Golang

**Original Framework**
(https://github.com/yosssi/gcss)

*Support.*
*[sass](http://sass-lang.com) and [stylus](http://stylus-lang.com)*
(css compression)

**Install**
```
$ go get -u github.com/tualeke/oh_sass
```

**Example usage**
```
	// source .sass and .styl file
	path , err := oh_sass.SassCompile("sass source")
	if err != nil {
		log.Println(err)
	}

	// output to your folder
	err = oh_sass.MoveTo(path,"output to your folder")
	if err != nil {
		log.Println(err)
```
**Example**
```
package main

import (
	"log"
	"github.com/tualeke/oh_sass"
)

func main() {
	// source .sass and .styl file
	path , err := oh_sass.SassCompile("template/static/sass/style.sass")
	if err != nil {
		log.Println(err)
	}

	// output to your folder
	err = oh_sass.MoveTo(path,"template/public/css/main.css")
	if err != nil {
		log.Println(err)
	}
}
```

##or##
```
package main


import (
	"log"
	"github.com/tualeke/oh_sass"
)

func main() {
	
	source := "template/static/sass/style.sass"  // source .sass and .styl file
	output := "template/public/css/main.css"     // output to your folder
	
	err := oh_sass.CompileTo(source,output)
	if err != nil {
		log.Println(err)
	}

}
```

##SYNTAX
.sass Example 
```
nav
  ul
    margin: 0
    padding: 0
    list-style: none

  li
    display: inline-block

  a
    display: block
    padding: 6px 12px
    text-decoration: none

```
compile to
```
nav ul {
  margin: 0;
  padding: 0;
  list-style: none;
}

nav li {
  display: inline-block;
}

nav a {
  display: block;
  padding: 6px 12px;
  text-decoration: none;
}

```

.styl
Example
```
border-radius()
  -webkit-border-radius: arguments
  -moz-border-radius: arguments
  border-radius: arguments

body a
  font: 12px/1.4 "Lucida Grande", Arial, sans-serif
  background: black
  color: #ccc

form input
  padding: 5px
  border: 1px solid
  border-radius: 5px
  
  ```
compile to

```
border-radius() {
  -webkit-border-radius: arguments
  -moz-border-radius: arguments
  border-radius: arguments
}

body a {
  font: 12px/1.4 "Lucida Grande", Arial, sans-serif;
  background: black;
  color: #ccc;
}

form input {
  padding: 5px;
  border: 1px solid;
  border-radius: 5px;
}
```

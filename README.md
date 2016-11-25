# oh_sass SASS and Stylus compile to css on Golang

**Original Framework**
(https://github.com/yosssi/gcss)

*Support.*
*[sass](http://sass-lang.com) and [stylus](http://stylus-lang.com)*
(css compression)

**Install**
```
$ go get github.com/tualeke/oh_sass
```

**Example usage**
```
	// source sass file
	path , err := oh_sass.SassCompile("sass source")
	if err != nil {
		log.Println(err)
	}

	// output to your folder
	err = oh_sass.MoveTo(path,"output to your folder")
	if err != nil {
		log.Println(err)
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

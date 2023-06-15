<h1>uconv</h1>
<h5>
  A simple CLI for doing unit conversions on the fly.
</h5>

<p>
  uconv allows you to convert values on the fly in your terminal, you can also add or remove your own conversions. </br>
  The project is open source written in Go, I learnt the language and features as I went. So it is definitely not perfect.
</p>

<p>
  At the moment the app stores the conversions in a textfile that gets saved to the documents folder. </br>
  The idea in the future is to use a different form of local storage to make it a bit easier. The idea at the moment is to use badger
</p>

<h3>
Planned updates:
</h3>
<ol>
<li>Improve installation files/instructions</li>
<li>Add Reset and Location commands</li>
<li>Change from text file to badger for conversion storage</li>
</ol>

<h3>
  Installation
</h3>
  <s>1. Download executable from the releases page, here</s>
  </br>
  2. Download the zip code and run it as a regular Go program </br>
  3. If you have go installed run
  
  `go install github.com/adams-Thomas/uconv@latest`

<h3>Usage</h3>
uconv comes with a set of conversions by default, they can be found here ADD LINK TO BOTTOM OF PAGE HERE</br></br>

<b>Base command usage for a conversion:</b>

`uconv [CONVERSION] [VALUE]`

Example:

`uconv rempx 1`, this will convert from rem to pixel

The command can be run in reverse:
`uconv pxrem 16`, this will convert from pixel to rem


<b>List available conversions</b>
The list command has a few options for formatting:
1. format: formats the list and is the default option:
```
1. From 1 cel to 32 fh
2. From 1 rem to 16 px
```
2. plain: displays as/is in the storage:
```
cel,fh,32
rem,px,16
```
3. usages: how each conversion gets used:
```
celfh
rempx
```


<b>Add a conversion</b>

`uconv add [FROM] [TO] [VALUE]`

Example:

`uconv add cm mm 10`
Only add the conversion if it does not exist in the storage.

<b>Remove a conversion</b>

`uconv remove [CONVERSION]`

Example:

`uconv remove rempx`
Will only delete the conversion if it exists in the storage


<h3>Initial Conversions</h3>
The list of conversions that come with the app, this can be subject to change.

```
cel,fh,32
rem,px,16
mile,km,1.60934
m,feet,3.28084
```
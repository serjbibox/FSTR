
<!doctype html>
<html>
  <head>
    <title>Simple Inventory API</title>
    <style type="text/css">
      body {
	font-family: Trebuchet MS, sans-serif;
	font-size: 15px;
	color: #444;
	margin-right: 24px;
}

h1	{
	font-size: 25px;
}
h2	{
	font-size: 20px;
}
h3	{
	font-size: 16px;
	font-weight: bold;
}
hr	{
	height: 1px;
	border: 0;
	color: #ddd;
	background-color: #ddd;
}

.app-desc {
  clear: both;
  margin-left: 20px;
}
.param-name {
  width: 100%;
}
.license-info {
  margin-left: 20px;
}

.license-url {
  margin-left: 20px;
}

.model {
  margin: 0 0 0px 20px;
}

.method {
  margin-left: 20px;
}

.method-notes	{
	margin: 10px 0 20px 0;
	font-size: 90%;
	color: #555;
}

pre {
  padding: 10px;
  margin-bottom: 2px;
}

.http-method {
 text-transform: uppercase;
}

pre.get {
  background-color: #0f6ab4;
}

pre.post {
  background-color: #10a54a;
}

pre.put {
  background-color: #c5862b;
}

pre.delete {
  background-color: #a41e22;
}

.huge	{
	color: #fff;
}

pre.example {
  background-color: #f3f3f3;
  padding: 10px;
  border: 1px solid #ddd;
}

code {
  white-space: pre;
}

.nickname {
  font-weight: bold;
}

.method-path {
  font-size: 1.5em;
  background-color: #0f6ab4;
}

.up {
  float:right;
}

.parameter {
  width: 500px;
}

.param {
  width: 500px;
  padding: 10px 0 0 20px;
  font-weight: bold;
}

.param-desc {
  width: 700px;
  padding: 0 0 0 20px;
  color: #777;
}

.param-type {
  font-style: italic;
}

.param-enum-header {
width: 700px;
padding: 0 0 0 60px;
color: #777;
font-weight: bold;
}

.param-enum {
width: 700px;
padding: 0 0 0 80px;
color: #777;
font-style: italic;
}

.field-label {
  padding: 0;
  margin: 0;
  clear: both;
}

.field-items	{
	padding: 0 0 15px 0;
	margin-bottom: 15px;
}

.return-type {
  clear: both;
  padding-bottom: 10px;
}

.param-header {
  font-weight: bold;
}

.method-tags {
  text-align: right;
}

.method-tag {
  background: none repeat scroll 0% 0% #24A600;
  border-radius: 3px;
  padding: 2px 10px;
  margin: 2px;
  color: #FFF;
  display: inline-block;
  text-decoration: none;
}

    </style>
  </head>
  <body>
  <h1>Simple Inventory API</h1>
    <div class="app-desc">This is a simple API</div>
    <div class="app-desc">More information: <a href="https://helloreverb.com">https://helloreverb.com</a></div>
    <div class="app-desc">Contact Info: <a href="you@your-company.com">you@your-company.com</a></div>
    <div class="app-desc">Version: 1.0.0</div>
    <div class="app-desc">BasePath:/serjbibox/API_FSTR/1.0.0</div>
    <div class="license-info">Apache 2.0</div>
    <div class="license-url">http://www.apache.org/licenses/LICENSE-2.0.html</div>
  <h2>Access</h2>

  <h2><a name="__Methods">Methods</a></h2>
  [ Jump to <a href="#__Models">Models</a> ]

  <h3>Table of Contents </h3>
  <div class="method-summary"></div>
  <h4><a href="#Admins">Admins</a></h4>
  <ul>
  <li><a href="#addInventory"><code><span class="http-method">post</span> /inventory</code></a></li>
  </ul>
  <h4><a href="#Developers">Developers</a></h4>
  <ul>
  <li><a href="#searchInventory"><code><span class="http-method">get</span> /inventory</code></a></li>
  </ul>

  <h1><a name="Admins">Admins</a></h1>
  <div class="method"><a name="addInventory"/>
    <div class="method-path">
    <a class="up" href="#__Methods">Up</a>
    <pre class="post"><code class="huge"><span class="http-method">post</span> /inventory</code></pre></div>
    <div class="method-summary">adds an inventory item (<span class="nickname">addInventory</span>)</div>
    <div class="method-notes">Adds an item to the system</div>


    <h3 class="field-label">Consumes</h3>
    This API call consumes the following media types via the <span class="header">Content-Type</span> request header:
    <ul>
      <li><code>application/json</code></li>
    </ul>

    <h3 class="field-label">Request body</h3>
    <div class="field-items">
      <div class="param">inventoryItem <a href="#InventoryItem">InventoryItem</a> (optional)</div>

      <div class="param-desc"><span class="param-type">Body Parameter</span> &mdash; Inventory item to add </div>

    </div>  <!-- field-items -->





    <!--Todo: process Response Object and its headers, schema, examples -->


    <h3 class="field-label">Produces</h3>
    This API call produces the following media types according to the <span class="header">Accept</span> request header;
    the media type will be conveyed by the <span class="header">Content-Type</span> response header.
    <ul>
      <li><code>application/json</code></li>
    </ul>

    <h3 class="field-label">Responses</h3>
    <h4 class="field-label">201</h4>
    item created
        <a href="#"></a>
    <h4 class="field-label">400</h4>
    invalid input, object invalid
        <a href="#"></a>
    <h4 class="field-label">409</h4>
    an existing item already exists
        <a href="#"></a>
  </div> <!-- method -->
  <hr/>
  <h1><a name="Developers">Developers</a></h1>
  <div class="method"><a name="searchInventory"/>
    <div class="method-path">
    <a class="up" href="#__Methods">Up</a>
    <pre class="get"><code class="huge"><span class="http-method">get</span> /inventory</code></pre></div>
    <div class="method-summary">searches inventory (<span class="nickname">searchInventory</span>)</div>
    <div class="method-notes">By passing in the appropriate options, you can search for
available inventory in the system</div>





    <h3 class="field-label">Query parameters</h3>
    <div class="field-items">
      <div class="param">searchString (optional)</div>

      <div class="param-desc"><span class="param-type">Query Parameter</span> &mdash; pass an optional search string for looking up inventory </div><div class="param">skip (optional)</div>

      <div class="param-desc"><span class="param-type">Query Parameter</span> &mdash; number of records to skip for pagination format: int32</div><div class="param">limit (optional)</div>

      <div class="param-desc"><span class="param-type">Query Parameter</span> &mdash; maximum number of records to return format: int32</div>
    </div>  <!-- field-items -->


    <h3 class="field-label">Return type</h3>
    <div class="return-type">
      array[<a href="#InventoryItem">InventoryItem</a>]
      
    </div>

    <!--Todo: process Response Object and its headers, schema, examples -->

    <h3 class="field-label">Example data</h3>
    <div class="example-data-content-type">Content-Type: application/json</div>
    <pre class="example"><code>{}</code></pre>

    <h3 class="field-label">Produces</h3>
    This API call produces the following media types according to the <span class="header">Accept</span> request header;
    the media type will be conveyed by the <span class="header">Content-Type</span> response header.
    <ul>
      <li><code>application/json</code></li>
    </ul>

    <h3 class="field-label">Responses</h3>
    <h4 class="field-label">200</h4>
    search results matching criteria
        
    <h4 class="field-label">400</h4>
    bad input parameter
        <a href="#"></a>
  </div> <!-- method -->
  <hr/>

  <h2><a name="__Models">Models</a></h2>
  [ Jump to <a href="#__Methods">Methods</a> ]

  <h3>Table of Contents</h3>
  <ol>
    <li><a href="#InventoryItem"><code>InventoryItem</code> - </a></li>
    <li><a href="#Manufacturer"><code>Manufacturer</code> - </a></li>
  </ol>

  <div class="model">
    <h3><a name="InventoryItem"><code>InventoryItem</code> - </a> <a class="up" href="#__Models">Up</a></h3>
    <div class='model-description'></div>
    <div class="field-items">
      <div class="param">id </div><div class="param-desc"><span class="param-type"><a href="#UUID">UUID</a></span>  format: uuid</div>
          <div class="param-desc"><span class="param-type">example: d290f1ee-6c54-4b01-90e6-d701748f0851</span></div>
<div class="param">name </div><div class="param-desc"><span class="param-type"><a href="#string">String</a></span>  </div>
          <div class="param-desc"><span class="param-type">example: Widget Adapter</span></div>
<div class="param">releaseDate </div><div class="param-desc"><span class="param-type"><a href="#DateTime">Date</a></span>  format: date-time</div>
          <div class="param-desc"><span class="param-type">example: 2016-08-29T09:12:33.001Z</span></div>
<div class="param">manufacturer </div><div class="param-desc"><span class="param-type"><a href="#Manufacturer">Manufacturer</a></span>  </div>
    </div>  <!-- field-items -->
  </div>
  <div class="model">
    <h3><a name="Manufacturer"><code>Manufacturer</code> - </a> <a class="up" href="#__Models">Up</a></h3>
    <div class='model-description'></div>
    <div class="field-items">
      <div class="param">name </div><div class="param-desc"><span class="param-type"><a href="#string">String</a></span>  </div>
          <div class="param-desc"><span class="param-type">example: ACME Corporation</span></div>
<div class="param">homePage (optional)</div><div class="param-desc"><span class="param-type"><a href="#string">String</a></span>  format: url</div>
          <div class="param-desc"><span class="param-type">example: https://www.acme-corp.com</span></div>
<div class="param">phone (optional)</div><div class="param-desc"><span class="param-type"><a href="#string">String</a></span>  </div>
          <div class="param-desc"><span class="param-type">example: 408-867-5309</span></div>
    </div>  <!-- field-items -->
  </div>
  </body>
</html>

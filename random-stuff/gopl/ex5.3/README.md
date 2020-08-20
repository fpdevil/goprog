# Exercuse 5.3

Write a function to print the content of all text nodes in an `HTML` document tree.
Do not descend into `<script>` or `<style>` elements, since their contents are not visible in a web browser.

## Some important information

- Node type constants

| Constant                           | Description                                                                |
|------------------------------------|----------------------------------------------------------------------------|
| `Node.ELEMENT_NODE`                | An Element node like <p> or <div>.                                         |
| `Node.TEXT_NODE`                   | The actual Text inside an Element or Attr.                                 |
| `Node.CDATA_SECTION_NODE`          | A CDATASection, such as <!CDATA[[ … ]]>.                                   |
| `Node.PROCESSING_INSTRUCTION_NODE` | A ProcessingInstruction of an XML document, such as <?xml-stylesheet … ?>. |
| `Node.COMMENT_NODE`                | A Comment node, such as <!-- … -->.                                        |
| `Node.DOCUMENT_NODE`               | A Document node.                                                           |
| `Node.DOCUMENT_TYPE_NODE`          | A DocumentType node, such as <!DOCTYPE html>.                              |
| `Node.DOCUMENT_FRAGMENT_NODE`      | A DocumentFragment node.                                                   |

- `TokenType` which will be used in tokenizing the `html` content is defined as below as per the documentation.

| `ErrorToken`          | ErrorToken means that an error occurred during tokenization. |
|-----------------------|--------------------------------------------------------------|
| `TextToken`           | TextToken means a text node.                                 |
| `StartTagToken`       | A StartTagToken looks like <a>.                              |
| `EndTagToken`         | An EndTagToken looks like </a>.                              |
| `SelfClosingTagToken` | A SelfClosingTagToken tag looks like <br/>.                  |
| `CommentToken`        | A CommentToken looks like <!--x-->.                          |
| `DoctypeToken`        | A DoctypeToken looks like <!DOCTYPE x>                       |
## Usage

```bash
⇒  go run main.go http://httpbin.org
* httpbin.org
* httpbin.org
* 0.9.2
* [ Base URL: httpbin.org/ ]
* A simple HTTP Request & Response Service.
* Run locally:
* $ docker run -p 80:80 kennethreitz/httpbin
* the developer - Website
* Send email to the developer
* [Powered by
* Flasgger
* ]
* Other Utilities
* HTML form
* that posts to /post /forms/post
```
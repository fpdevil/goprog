# Exercise 5.1 HTML Parsing

Change the `findlinks` program to traverse the `n.FirstChild` linked list using recursive calls to visit instead of a loop.

## Some important information

- Node type constants

| Constant                           | Description                                                                |
|------------------------------------|----------------------------------------------------------------------------|
| `Node.ELEMENT_NODE`                | An Element node like `<p>` or `<div>`.                                     |
| `Node.TEXT_NODE`                   | The actual Text inside an Element or Attr.                                 |
| `Node.CDATA_SECTION_NODE`          | A CDATASection, such as <!CDATA[[ … ]]>.                                   |
| `Node.PROCESSING_INSTRUCTION_NODE` | A ProcessingInstruction of an XML document, such as <?xml-stylesheet … ?>. |
| `Node.COMMENT_NODE`                | A Comment node, such as `<!-- … -->`.                                      |
| `Node.DOCUMENT_NODE`               | A Document node.                                                           |
| `Node.DOCUMENT_TYPE_NODE`          | A DocumentType node, such as `<!DOCTYPE html>`.                            |
| `Node.DOCUMENT_FRAGMENT_NODE`      | A DocumentFragment node.                                                   |

- `TokenType` which will be used in tokenizing the `html` content is defined as below as per the documentation.

| `ErrorToken`          | ErrorToken means that an error occurred during tokenization. |
|-----------------------|--------------------------------------------------------------|
| `TextToken`           | TextToken means a text node.                                 |
| `StartTagToken`       | A StartTagToken looks like `<a>`.                            |
| `EndTagToken`         | An EndTagToken looks like `</a>`.                            |
| `SelfClosingTagToken` | A SelfClosingTagToken tag looks like `<br/>`.                |
| `CommentToken`        | A CommentToken looks like `<!--x-->`.                        |
| `DoctypeToken`        | A DoctypeToken looks like `<!DOCTYPE x>`                     |

## Usage

```bash
⇒  go run main.go https://golang.org
  0: https://support.eji.org/give/153413/#!/donation/checkout
  1: /
  2: /doc/
  3: /pkg/
  4: /project/
  5: /help/
  6: /blog/
  7: https://play.golang.org/
  8: /dl/
  9: https://tour.golang.org/
 10: https://blog.golang.org/
 11: /doc/copyright.html
 12: /doc/tos.html
 13: http://www.google.com/intl/en/policies/privacy/
 14: http://golang.org/issues/new?title=x/website:
 15: https://google.com
```

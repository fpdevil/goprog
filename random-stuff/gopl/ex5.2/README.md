# Exercuse 5.2

Write a function to populate a mapping from element names

`p`, `div`, `span` and so on - to the number of elements with that name in an HTML document tree

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
2     title
4     svg
6     section
4     h2
2     html
1     meta
14    symbol
32    div
4     pre
8     script
2     body
14    path
10    a
2     defs
2     hgroup
5     br
2     b
3     link
2     style
2     p
2     code
2     span
2     ul
2     li
2     head
2     small
```
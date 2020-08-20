# Exercuse 5.5

Implement the `countWordsAndImages` function.

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
⇒  go run main.go http://gopl.io
words: 198
images: 4

⇒  go run main.go http://golang.org
words: 123
images: 3
```
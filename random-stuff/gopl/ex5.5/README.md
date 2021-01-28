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

### Testing with multiple url's at once
```bash
⇒  go run main.go https://golang.org http://gopl.io https://httpbin.org http://a.b.c.d
------------------------------------------------
[url: https://httpbin.org]
------------------------------------------------
# words   : 45
# images  : 0
------------------------------------------------
[url: http://a.b.c.d]
------------------------------------------------
# error   : http GET error http://a.b.c.d: Get "http://a.b.c.d": dial tcp: lookup a.b.c.d: no such host
# words   : 0
# images  : 0
------------------------------------------------
[url: https://golang.org]
------------------------------------------------
# words   : 123
# images  : 3
------------------------------------------------
[url: http://gopl.io]
------------------------------------------------
# words   : 198
# images  : 4
```

### Testing with an invalid url
```bash
⇒  go run main.go ftp://1.2.3.4
------------------------------------------------
[url: ftp://1.2.3.4]
------------------------------------------------
# error   : http GET error ftp://1.2.3.4: Get "ftp://1.2.3.4": unsupported protocol scheme "ftp"
# words   : 0
# images  : 0
```
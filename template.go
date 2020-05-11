package main

const indexTpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .Title }}</title>
</head>
<body>
    {{ .Body }}
</body>
</html>
`
const blank = "Please tell me your favorite tree"
const treePrefix = "It's nice to know that your favorite tree is a "

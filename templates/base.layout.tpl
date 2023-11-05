<!-- vim: set ft=html: -->
{{define "base"}}
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title>Home</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
        <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL" crossorigin="anonymous"></script>
    </head>
    <body>
        <div class="container">
            <div class="row">
                <div class="content">
                    {{with .Flash}}
                    <div class="mt-3 alert alert-success" role="alert">
                        {{.}}
                    </div>
                    {{end}}
                    {{with .Error}}
                    <div class="mt-3 alert alert-danger" role="alert">
                        {{.}}
                    </div>
                    {{end}}
                </div>
            </div>
        </div>
        {{block "content" .}}
        {{end}}
    </body>
</html>
{{end}}

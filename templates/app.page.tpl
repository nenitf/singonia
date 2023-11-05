<!-- vim: set ft=html: -->
{{template "base" .}}

{{define "content"}}
<div class="container">
    <div class="row">
        <div class="col">
            <small>Logged as {{.User.Name}}</small>
        </div>
    </div>
</div>
{{end}}

<!-- vim: set ft=html: -->
{{template "base" .}}

{{define "content"}}
<div class="container">
    <div class="row">
        <div class="col">
            <h1 class="mt-3">Singonia</h1>
            <form action="" method="post">
                <div class="form-group">
                    <label for="name">Qual seu nome?</label>
                    <input type="text" class="form-control" id="name" name="name">
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
        </div>
    </div>
</div>
{{end}}

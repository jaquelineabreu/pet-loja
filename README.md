## Aplicação web em Go - Pets Loja

**_Documentação_**

_Arquivos .html_

- Para melhorar a aparência do template utilizei o framework bootstrap via CDN

_main.go_

- Ouve e responde requisições - func http.ListenAndServe(addr string, handler http.Handler) error
- Retorna os templates - func Must(t *Template, err error) *Template
- Rotas (requisição/resposta) - http.HandleFunc(pattern string, handler func(http.ResponseWriter, \*http.Request))

_Obs:_ :shipit:

- Para executar o index.html foi embedado código go no inicio e no final da mesma.

  Ex:
  {{define "Index"}}
  <bloco html>
  {{end}}

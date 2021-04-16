---
Title: Get Vim Ready for Go Programming
Query: true
---

The [`fatih/vim-go`](https://github.com/fatih/vim-go) plugin is pretty standard albeit annoying at times. Don't forget to Install all the `vim-go` dependencies with `:GoInstallBinaries` from within a Vim session if you choose to use the following:

```vim
" Install the Plug plugin manager if not detected.
if empty(glob('~/.vim/autoload/plug.vim'))
  silent !curl -fLo ~/.vim/autoload/plug.vim --create-dirs
    \ https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
  autocmd VimEnter * PlugInstall
endif

" If Plug plugin manager detected then load the plugins and configure
if filereadable(expand("~/.vim/autoload/plug.vim"))
  call plug#begin('~/.vimplugins')
  " ... other plugins here
  Plug 'fatih/vim-go'
  call plug#end()
  let g:go_fmt_fail_silently = 0     " let me out even with errors
  let g:go_fmt_command = 'goimports' " autoupdate import
  let g:go_fmt_autosave = 1          " autosave on updates

" Otherwise fallback to some safe backups
else
  autocmd vimleavepre *.go !gofmt -w % " backup if fatih fails
endif

```

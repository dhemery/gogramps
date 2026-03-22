local config = { settings = { gopls = {} } }
config.settings.gopls['local'] = 'github.com/dhemery/gogramps'
vim.lsp.config('gopls', config)

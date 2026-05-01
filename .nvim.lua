local dap = require("dap")

vim.api.nvim_create_autocmd("VimEnter", {
	once = true,
	callback = function()
		vim.list_extend(dap.configurations.go, {
			{
				type = "go",
				name = "Debug gonotepad (attach)",
				request = "attach",
				mode = "remote",
				host = "127.0.0.1",
				port = 2345,
				initialize = function()
					vim.fn.system("pkill dlv")
					vim.fn.jobstart({
						"alacritty",
						"-e",
						"dlv",
						"debug",
						"--headless",
						"--listen=:2345",
						"--accept-multiclient",
						".",
						"--",
						"editor/testdata/simple.txt",
					}, { detach = true })
					vim.uv.sleep(2000) -- block until delve is up
				end,
			},
		})
	end,
})

vim.keymap.set("n", "<leader>Dx", function()
	dap.terminate()
	vim.fn.system("pkill dlv")
	require("dapui").close()
end, { desc = "Debug: terminate session" })

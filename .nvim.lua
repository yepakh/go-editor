local dap = require("dap")

-- dap.configurations.go = {
-- 	{
-- 		type = "go",
-- 		name = "Debug gonotepad (attach)",
-- 		request = "attach",
-- 		mode = "remote",
-- 		host = "127.0.0.1",
-- 		port = 2345,
-- 	},
-- }

vim.keymap.set("n", "<leader>Dd", function()
	-- kill any previous delve session
	vim.fn.system("pkill dlv")

	-- spawn delve headless in a new alacritty window
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
		"go.sum",
	}, { detach = true })

	-- wait for delve to start, then attach
	vim.defer_fn(function()
		dap.continue()
	end, 2000)
end, { desc = "Debug: launch + attach gonotepad" })

vim.keymap.set("n", "<leader>Dx", function()
	dap.terminate()
	vim.fn.system("pkill dlv")
	require("dapui").close()
end, { desc = "Debug: terminate session" })

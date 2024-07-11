function toggleDarkMode() {
	document.documentElement.classList.toggle("dark")
	const currentTheme = document.documentElement.getAttribute("data-theme")
	if (!currentTheme ||
		currentTheme === "corporate") {
		document.documentElement.setAttribute("data-theme", "dim")
	} else {
		document.documentElement.setAttribute("data-theme", "corporate")
	}
	localStorage.setItem("theme", document.documentElement.getAttribute("data-theme"))
}

function loadTheme() {
	const currentTheme = localStorage.getItem("theme") ?? "dim"
	console.log(currentTheme)
	document.documentElement.setAttribute(
		"data-theme",
		currentTheme
	)
}
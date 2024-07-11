function toggleDarkMode() {
	document.documentElement.classList.toggle("dark")
	const currentTheme = document.documentElement.getAttribute("data-theme")
	if (currentTheme &&
		currentTheme === "corporate") {
		document.documentElement.setAttribute("data-theme", "dim")
	} else {
		document.documentElement.setAttribute("data-theme", "corporate")
	}
}
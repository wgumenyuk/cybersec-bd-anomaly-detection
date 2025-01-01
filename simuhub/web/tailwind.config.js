/**
	@type {import("tailwindcss").Config}
*/
export default {
	content: [
		"index.html",
		"src/**/*.{css,js,ts}"
	],
	theme: {
		extend: {
			colors: {
				silver: {
					50: "#F2F2F3",
					100: "#E3E2E4",
					200: "#C9C8CB",
					300: "#ACACAF",
					400: "#939296",
					500: "#77767B",
					600: "#605F63",
					700: "#474649",
					800: "#302F31",
					900: "#171617",
					950: "#0D0C0D"
				}
			},
			fontFamily: {
				sans: "Roboto, sans-serif"
			}
		}
	},
	plugins: []
};

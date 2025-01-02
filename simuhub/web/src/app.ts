import Alpine from "alpinejs";
import { createIcons, Check, OctagonAlert } from "lucide";

// Styles
import "@fontsource/roboto/latin-400.css";
import "./app.css";

type Weights = Record<string, {
	label: string;
	value: number;
}>;

Alpine.data("status", () => ({
	connected: false,
	async init() {
		try {
			const response = await fetch("/api/v1/status");
			this.connected = !import.meta.env.DEV && response.ok;
		} catch {
			this.connected = false;
		}
	}
}));

Alpine.data("config", () => ({
	t: 5,
	weights: {
		normal: {
			label: "Normal Traffic",
			value: 100
		},
		bruteforce: {
			label: "Bruteforce",
			value: 0
		},
		ddos: {
			label: "DDoS",
			value: 0
		},
	} as Weights,
	ok: false,
	async init() {
		const response = await fetch("/api/v1/config", {
			method: "GET"
		});

		if (!response.ok) {
			return;
		}

		const config = await response.json() as Record<string, number>;

		this.t = config.t;

		Object
			.entries(config)
			.filter(([key]) => key !== "t")
			.forEach(([key, value]) => {
				this.weights[key].value = value * 100;
			});
	},
	get total() {
		return Object
			.values(this.weights)
			.reduce((sum, weight) => sum + weight.value, 0);
	},
	remaining(changedKey: string) {
		const sum = Object
			.entries(this.weights)
			.reduce((sum, [key, weight]) => {
				return (key !== changedKey) ? sum + weight.value : sum;
			}, 0);

		return 100 - sum;
	},
	reset() {
		this.t = 5;

		Object.keys(this.weights).forEach((key) => {
			this.weights[key].value = (key === "normal") ? 100 : 0;
		});
	},
	async update() {
		if (this.total !== 100 || this.ok) {
			return;
		}

		const body = Object
			.entries(this.weights)
			.reduce((body: Record<string, any>, [key, weight]) => {
				body[key] = weight.value / 100;
				return body;
			}, {});

		body.t = Number(this.t);

		const response = await fetch("/api/v1/config", {
			method: "PUT",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify(body)
		});

		if (!response.ok) {
			console.error(response);
			return;
		}

		this.ok = true;

		setTimeout(() => {
			this.ok = false;
		}, 3000);
	}
}));

createIcons({
	icons: {
		Check,
		OctagonAlert
	}
});

Alpine.start();

import { selector } from "recoil";
import * as types from "@/proto/types";

export const AppMeta = selector<types.AppMeta>({
	key: "AppMeta",
	get: async () => {
		try {
			const response = await fetch("/api/getMeta");
			const data = await response.json();
			return data;
		} catch (error) {}
	},
});

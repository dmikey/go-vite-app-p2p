import { RecoilRoot } from "recoil";
// import { PlaygroundPage } from "./playground";
import { AuthenticationPage } from "./authentication";
// import { SettingsPage } from "./settings";
// import { SettingsProfilePage } from "@/components/ui/settings-profile";
// import { TabsLayout } from "./Layout";
// import { DashboardPage } from "./dashboard";
import "./App.css";

function App() {
	return (
		<>
			<RecoilRoot>
				{/* <DashboardPage /> */}
				{/* <SettingsPage>
				<SettingsProfilePage />
				</SettingsPage> */}
				<AuthenticationPage />
				{/* <PlaygroundPage /> */}
			</RecoilRoot>
		</>
	);
}

export default App;

import { Card, CardContent } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Separator } from "@/components/ui/separator";
import { PlaygroundPage } from "./playground";
import { SettingsPage } from "./settings";
import { SettingsProfilePage } from "@/components/ui/settings-profile";

export function TabsLayout() {
	return (
		<div className="hidden h-full flex-col md:flex">
			<div className="container flex flex-col items-start justify-between space-y-2 py-4 sm:flex-row sm:items-center sm:space-y-0 md:h-16">
				<img src="/vite.svg" alt="myapp" width="38px" />
				<h2 className="text-lg font-semibold">&nbsp;myapp</h2>
				<div className="ml-auto flex w-full space-x-2 sm:justify-end" />
			</div>
			<Separator />
			<Tabs defaultValue="account" className="flex-1">
				<TabsList className="grid w-full grid-cols-2">
					<TabsTrigger value="account">Playground</TabsTrigger>
					<TabsTrigger value="settings">Settings</TabsTrigger>
				</TabsList>
				<TabsContent value="account">
					<Card>
						<CardContent className="space-y-2">
							<PlaygroundPage />
						</CardContent>
					</Card>
				</TabsContent>
				<TabsContent value="settings">
					<Card>
						<CardContent className="space-y-2">
							<CardContent className="space-y-2">
								<SettingsPage>
									<SettingsProfilePage />
								</SettingsPage>
							</CardContent>
						</CardContent>
					</Card>
				</TabsContent>
			</Tabs>
		</div>
	);
}

import { writable } from "svelte/store";
import { main } from "./wailsjs/go/models";

export enum View {
    Main,
    Config,
    PunchOut
}

export const issue = writable(new main.CurrentIssue())
export const config = writable(new main.Config())
export const view = writable(View.Main)
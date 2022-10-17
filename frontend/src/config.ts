import type { main } from "./wailsjs/go/models";

export function configComplete(c: main.Config): boolean {
    return c.Token !== "" && c.Server !== ""
}
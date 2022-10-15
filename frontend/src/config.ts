import type { main } from "./wailsjs/go/models";

export function configComplete(c: main.Config): boolean {
    return c.Username !== "" && c.Password !== "" && c.Server !== ""
}
export interface ITokenService {
    remove(): void;
    add(token: string): void
    get(): string | null
}

export class TokenService implements ITokenService {
    private alise: string = "token"
    constructor() {

    }
    get(): string | null {
        return localStorage.getItem(this.alise)
    }
    remove(): void {
        localStorage.removeItem(this.alise);
    }
    add(token: string): void {
        localStorage.setItem(this.alise, token)
    }
}

export const defaultTokenService:ITokenService = new TokenService();
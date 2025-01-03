
import { defaultTokenService } from '@/service/TokenService';
import { Mgr } from './oidcConfig'


Mgr.signinRedirectCallback().then(() => {
    Mgr.getUser().then((u) => {
        if (u) {
            defaultTokenService.add(u.token_type + " " + u.access_token);
            window.location.href = "/";
        } else {
            Mgr.signinRedirect();
        }
    })
})
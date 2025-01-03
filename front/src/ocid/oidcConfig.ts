import { defaultTokenService } from '@/service/TokenService';
import { UserManager } from 'oidc-client'

const BaseUrl = window.location.origin;

export const Mgr = new UserManager({
    authority: BaseUrl,
    client_id: "interactive",
    redirect_uri: `${BaseUrl}/signin-oidc.html`,
    response_type: "code",
    scope: "openid profile",
    post_logout_redirect_uri: `${BaseUrl}/signout-callback-oidc`,
    // stateStore: new CustomStateStore(sessionStorage)
    // : sessionStorage  // 使用 sessionStorage 存储 state
});


// 监听用户注销事件
Mgr.events.addUserSignedOut(function () {
    // localStorage.removeItem("token");
    defaultTokenService.remove();

    console.log('User has signed out.');
    Mgr.signoutRedirect();
});

// 监听用户签名事件（用户会话过期或清除）
Mgr.events.addUserUnloaded(function () {
    defaultTokenService.remove();

    console.log('User session is unloaded.');
    Mgr.signoutRedirect();    // 在这里执行用户会话被清除后的操作
});

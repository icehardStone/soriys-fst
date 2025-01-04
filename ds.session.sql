SELECT * from "ClientRedirectUris";

-- insert into "ClientRedirectUris"


-- INSERT INTO "ClientRedirectUris" ( "RedirectUri", "ClientId")
-- VALUES (
--     'http://192.168.137.1:3000/signin-oidc.html',
--     2
--   );


delete from "ClientRedirectUris" where "Id" in (6,7,8)

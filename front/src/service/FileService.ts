import { httpClient } from "./HttpClientService"
import { defaultTokenService } from "./TokenService"

export class BaseEntity {
    id!: number
}
export class FileModel extends BaseEntity {
    createTime?: Date
    name!: string
    size!: bigint
    tags?: string[]
    type?: string

    assign(o: object): FileModel {
        Object.assign(this, o)
        return this;
    }
}
// 下载文件接口定义
interface IFileService {
    remove(id: number): Promise<void>;
    getDownloadUrl(id: number): string;
    download(id: number, name: string): void;
    getList(pageSize?: number, pageIndex?: number): Promise<FileModel[]>;
}
// 实际实现
class FileService implements IFileService {
    async remove(id: number): Promise<void> {
        const resp = await httpClient.delete(`/api/file/${id}`);
        if (resp.status != 200) {
            console.log(resp.data)
        }
    }
    getDownloadUrl(id:number): string {
        const token = defaultTokenService.get() ?? "";
        // return window.location.origin + `/api/file/${id}/download?token=` + encodeURIComponent(token);
         return `/api/file/${id}/download?token=` + encodeURIComponent(token);
    }
    download(id: number, name: string): void {
       
        const href = this.getDownloadUrl(id);

        // console.log(href);
        var a = document.createElement('a');
        a.target = "_blank";
        a.href = href;
        a.download = name;
        a.click();

        // window.location.href = href;
        // window.open(href,"_blank");
        
    }
    async getList(pageSize?: number, pageIndex?: number): Promise<FileModel[]> {
        const reps = await httpClient.get('/api/file', {
            params: {
                pageSize: pageSize ?? 100,
                pageIndex: pageIndex ?? 1,
                key: null
            }
        });
        // debugger;
        var list = reps.data.data
        if (!list) {
            return [];
        }
        const cys = list.map((s: object) => new FileModel().assign(s));
        return cys;
    }
}
const fileService: IFileService = new FileService();
export {
    fileService
}
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
    download(id: number, name: string): void {
        var a = document.createElement('a')
        var event = new MouseEvent('click')

        a.download = name;
        a.href = `/api/file/${id}/download?token=${defaultTokenService.get()}`;

        a.dispatchEvent(event)
        // document.removeChild(a);
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
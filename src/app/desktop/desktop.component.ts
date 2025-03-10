import { Component } from '@angular/core';
import { Router } from "@angular/router";
import { RequestService } from "../request.service";
import { NzModalRef, NzModalService } from "ng-zorro-antd/modal";
import { WindowComponent } from "../window/window.component";
import { OemService } from "../oem.service";
import { AppService } from "../app.service";
import { UserService } from "../user.service";
import { PasswordComponent } from "../user/password/password.component"
declare var window: any;

@Component({
  selector: 'app-desktop',
  templateUrl: './desktop.component.html',
  styleUrls: ['./desktop.component.scss']
})
export class DesktopComponent {

  userInfo: any;
  constructor(
    private router: Router,
    private rs: RequestService,
    private ms: NzModalService,
    private us: UserService,
    protected os: OemService,
    protected _as: AppService) {
    this.userInfo = us && us.user;
  }
  handlePassword() {
    const modal: NzModalRef = this.ms.create({
      nzTitle: '修改密码',
      nzCentered: true,
      nzContent: PasswordComponent,
      nzFooter: [
        {
          label: '取消',
          onClick: () => {
            modal.destroy();
          }
        },
        {
          label: '保存',
          type: 'primary',
          onClick: componentInstance => {
            componentInstance!.submit().then(() => {
              modal.destroy();
            }, () => { })
          }
        }
      ]
    });
  }
  open(app: any) {
    if (window.innerWidth < 800) {
      this.router.navigate([app.entries[0].path])
      return;
    }

    this.ms.create({
      nzTitle: app.name,
      nzFooter: null,
      //nzMask: false,
      nzMaskClosable: false,
      nzWidth: "90%",
      //nzStyle: {height: "90%"},
      nzBodyStyle: { padding: "0", overflow: "hidden" },
      nzContent: WindowComponent,
      nzComponentParams: {
        entries: app.entries || []
      }
    })
  }

  logout() {
    this.rs.get("logout").subscribe(res => { }).add(() => this.router.navigateByUrl("/login"))
  }
}

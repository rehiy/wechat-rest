import { Component } from '@angular/core';

import { LevelData } from 'src/openapi/const';
import { RobotApi, ChatroomFetchAllParam, TablesChatroom } from '../../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../../openapi/wcfrest';


@Component({
    selector: 'page-bot-chatrooms',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class BotChatroomsComponent {

    public levels = LevelData;

    public contacts: Record<string, WcfrestContactPayload> = {};

    public chatrooms: Array<TablesChatroom> = [];

    constructor() {
        this.getContacts();
        this.getChatrooms();
    }

    public getContacts() {
        WrestApi.contacts().then((data) => {
            data.forEach((item) => this.contacts[item.wxid] = item);
        });
    }

    public getChatrooms() {
        const rq = {} as ChatroomFetchAllParam;
        RobotApi.chatroomList(rq).then((data) => {
            this.chatrooms = data || [];
        });
    }

}
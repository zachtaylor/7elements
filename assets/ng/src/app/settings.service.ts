import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { ChatSetting, QueueSetting } from './api';

@Injectable({
  providedIn: 'root'
})
export class SettingsService {
  // Runtime settings; TODO load cookies
  chat$ = new BehaviorSubject<ChatSetting>(new ChatSetting())
  queue$ = new BehaviorSubject<QueueSetting>(new QueueSetting())

  constructor() { }
}

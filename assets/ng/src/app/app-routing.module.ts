import { NgModule } from '@angular/core'
import { Routes, RouterModule } from '@angular/router'
import { IntroComponent } from './intro/intro.component'
import { MyAccountComponent } from './myaccount/myaccount.component'
import { CardsComponent } from './cards/cards.component'
import { DecksComponent } from './decks/decks.component'
import { DecksIdComponent } from './decks.id/decks.id.component'
import { IndexComponent } from './index/index.component'
import { ChatsComponent } from './chats/chats.component'
import { PlayComponent } from './play/play.component'
import { UpdatesComponent } from './updates/updates.component'
import { MyDecksIdComponent } from './mydecks.id/mydecks.id.component'
import { LostComponent } from './lost/lost.component'
import { TodoComponent } from './todo/todo.component'
import { HowtoComponent } from './howto/howto.component'
import { PacksComponent } from './packs/packs.component'
import { AdminSinkComponent } from './admin-sink/admin-sink.component'

const routes: Routes = [
  { path: 'admin/sink', component: AdminSinkComponent },
  { path: 'intro', component: IntroComponent },
  { path: 'myaccount', component: MyAccountComponent },
  { path: 'cards', component: CardsComponent },
  { path: 'packs', component: PacksComponent },
  { path: 'decks', component: DecksComponent },
  { path: 'decks/:id', component: DecksIdComponent },
  { path: 'mydecks/:id', component: MyDecksIdComponent },
  { path: 'chats', component: ChatsComponent },
  { path: 'play', component: PlayComponent },
  { path: 'updates', component: UpdatesComponent },
  { path: 'lost', component: LostComponent },
  { path: 'todo', component: TodoComponent },
  { path: 'how-to', component: HowtoComponent },
  { path: '', component: IndexComponent },
  {
    path: '*',
    redirectTo: '/',
    pathMatch: 'full'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core'
import { BrowserModule } from '@angular/platform-browser'
import { HttpClientModule } from '@angular/common/http'
import { CookieService } from 'ngx-cookie-service'
import { AppRoutingModule } from './app-routing.module'
import { AppComponent } from './app.component'
import { ArticleComponent } from './article/article.component'
import { IntroComponent } from './intro/intro.component'
import { MyAccountComponent } from './myaccount/myaccount.component'
import { CardDetailComponent } from './card-detail/card-detail.component'
import { IconComponent } from './icon/icon.component'
import { EscapeHtmlPipe, MapKeysPipe, MapValuesPipe, CountPipe } from './app.pipes'
import { DeckSummaryComponent } from './deck-summary/deck-summary.component'
import { DecksIdComponent } from './decks.id/decks.id.component'
import { CardsComponent } from './cards/cards.component'
import { IndexComponent } from './index/index.component'
import { ChatsComponent } from './chats/chats.component'
import { PlayComponent } from './play/play.component'
import { GameSeatComponent } from './game-seat/game-seat.component'
import { GameTokenComponent } from './game-token/game-token.component'
import { FormsModule, ReactiveFormsModule } from '@angular/forms'
import { CarouselComponent } from './carousel/carousel.component'
import { ToggleComponent } from './toggle/toggle.component'
import { UpdatesComponent } from './updates/updates.component'
import { MyDecksIdComponent } from './mydecks.id/mydecks.id.component'
import { LostComponent } from './lost/lost.component'
import { TodoComponent } from './todo/todo.component'
import { HowtoComponent } from './howto/howto.component'
import { PacksComponent } from './packs/packs.component'
import { DecksComponent } from './decks/decks.component'
import { QueueFormComponent } from './queue-form/queue-form.component'
import { PlayStackViewerComponent } from './play-stack-viewer/play-stack-viewer.component'
import { PlayChoiceViewerComponent } from './play-choice-viewer/play-choice-viewer.component';
import { AdminSinkComponent } from './admin-sink/admin-sink.component';
import { FormDetailComponent } from './form-detail/form-detail.component';
import { NavComponent } from './nav/nav.component';
import { PlayHandComponent } from './play-hand/play-hand.component'

@NgModule({
  declarations: [
    AppComponent,
    IntroComponent,
    MyAccountComponent,
    ArticleComponent,
    CardDetailComponent,
    IconComponent,
    EscapeHtmlPipe,
    MapKeysPipe,
    MapValuesPipe,
    CountPipe,
    DeckSummaryComponent,
    DecksIdComponent,
    CardsComponent,
    IndexComponent,
    ChatsComponent,
    PlayComponent,
    GameSeatComponent,
    GameTokenComponent,
    CarouselComponent,
    ToggleComponent,
    UpdatesComponent,
    MyDecksIdComponent,
    LostComponent,
    TodoComponent,
    HowtoComponent,
    PacksComponent,
    DecksComponent,
    QueueFormComponent,
    PlayStackViewerComponent,
    PlayChoiceViewerComponent,
    AdminSinkComponent,
    FormDetailComponent,
    NavComponent,
    PlayHandComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [CookieService],
  bootstrap: [AppComponent],
  schemas: [NO_ERRORS_SCHEMA]
})
export class AppModule { }

import { NgModule } from '@angular/core';
import { IonicApp, IonicModule } from 'ionic-angular';
import { MyApp } from './app.component';
import { HomePage } from '../pages/home/home';
import { TabsPage } from '../pages/tabs/tabs';
import {CameraPage} from "../pages/camera/camera";
import {HistoryPage} from "../pages/history/history";
import {NoticePage} from "../pages/notice/notice";
import {SearchList} from "../pages/search-list/search-list";
import {ChemDetail} from "../pages/chem-detail/chem-detail";
import {ChemSearchService} from "../providers/chem-search-service";

@NgModule({
  declarations: [
    MyApp,
    HomePage,
    TabsPage,
    CameraPage,
    HistoryPage,
    NoticePage,
    SearchList,
    ChemDetail
  ],
  imports: [
    IonicModule.forRoot(MyApp)
  ],
  bootstrap: [IonicApp],
  entryComponents: [
    MyApp,
    HomePage,
    TabsPage,
    CameraPage,
    HistoryPage,
    NoticePage,
    SearchList,
    ChemDetail
  ],
  providers: [ChemSearchService]
})
export class AppModule {}

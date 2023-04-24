import { Block, theme } from "galio-framework";
import { StyleSheet, Dimensions, ScrollView } from "react-native";
import EventItem from "../../components/EventItem";
import { useState } from "react";
const { width } = Dimensions.get("screen");
import { eventsList } from "../../constants/ExploreTab/bookmarkedEvents";
/*
View Bookmarks page allows user to view all their saved events from their profile. When they tap on the bookmark icon, 
they will be navigated to their bookmarks where a list of saved event items are displayed 
 */
const ViewBookmarks = () => {
  //state to handle event list, intial eventsList is grabbed from a constants folder called bookmarkedEvents
  const [events, setEvents] = useState(eventsList);

  /*
  handles removing an event if a user chooses to do so. This handler is passed 
  to each event item. If a user chooses to remove an event, the event id is passed 
  back to this handler. It then updates the state of the events list. 

  parameters
  idToRemove - Type (int):  event id that we want to remove from the list of events 
   */
  const handleRemoveItem = (idToRemove) => {
    setEvents(events.filter((item) => item.id !== idToRemove));
  };

  //creating Event items and passing props
  const allEvents = events.map(function (data, index) {
    return (
      <EventItem key={index} data={data} handleRemoveItem={handleRemoveItem} />
    );
  });

  return (
    <Block flex center style={styles.home}>
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.articles}
      >
        <Block flex center>
          {allEvents}
        </Block>
      </ScrollView>
    </Block>
  );
};

export default ViewBookmarks;

const styles = StyleSheet.create({
  home: {
    width: width,
  },
  articles: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

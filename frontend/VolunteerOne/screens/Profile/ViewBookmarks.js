import { Text, Block, theme } from "galio-framework";
import { StyleSheet, Dimensions, ScrollView } from "react-native";
import EventItem from "../../components/EventItem";
import { useState, useEffect } from "react";
const { width } = Dimensions.get("screen");
import { eventsList } from "../../constants/ExploreTab/bookmarkedEvents";

const ViewBookmarks = () => {
  const [events, setEvents] = useState(eventsList);

  const handleRemoveItem = (idToRemove) => {
    console.log("Id to remove: ", idToRemove);
    setEvents(events.filter((item) => item.id !== idToRemove));
  };

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

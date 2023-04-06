import EventDetails from "../../components/EventDetails";
import { Block, theme } from "galio-framework";
import { StyleSheet, Dimensions, ScrollView } from "react-native";
const { width } = Dimensions.get("screen");

/*
Description:
  This screen shows the event details of an organization's post when a user clicks on "View Event".
Props received:
  eventID - receives the eventID of a post, and then passes it to the EventDetails component 
            so that it can render the details of that event
*/

const Event = ({ route }) => {
  const eventID = route.params.eventID;
  return (
    <Block flex center style={styles.home}>
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.articles}
      >
        <Block flex center>
          <EventDetails eventID={eventID} />
        </Block>
      </ScrollView>
    </Block>
  );
};

const styles = StyleSheet.create({
  home: {
    width: width,
  },
  articles: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});
export default Event;

import EventPost from "../../components/EventPost";
import { Block, theme } from "galio-framework";
import { StyleSheet, Dimensions, ScrollView } from "react-native";
const { width } = Dimensions.get("screen");
const Event = ({ route }) => {
  const eventID = route.params.eventID;
  return (
    <Block flex center style={styles.home}>
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.articles}
      >
        <Block flex center>
          <EventPost eventID={eventID} />
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

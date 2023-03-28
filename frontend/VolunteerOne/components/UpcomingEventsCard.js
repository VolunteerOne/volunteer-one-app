import { StyleSheet } from "react-native";
import { Block, Text } from "galio-framework";
import { Image, TouchableWithoutFeedback } from "react-native";

/*
Description:
  This component returns an upcoming events card. Receives the data to insert into the card. 
Props received:
  data: date, title, image
*/

const UpcomingEventsCard = ({ navigation, data }) => {
  const { date, title, image } = data;
  return (
    <TouchableWithoutFeedback onPress={() => navigation.navigate("Pro")}>
      <Block row style={[styles.card, styles.shadowProp]}>
        <Block flex style={styles.imageContainer}>
          <Image source={{ uri: image }} style={styles.image} />
        </Block>
        <Block flex={2} style={styles.cardContent}>
          <Text size={14} color="#525F7F" style={styles.textSpacing}>
            {title}
          </Text>
          <Text size={12} color="#525F7F" style={styles.textSpacing}>
            on {date}
          </Text>
        </Block>
      </Block>
    </TouchableWithoutFeedback>
  );
};

const styles = StyleSheet.create({
  textSpacing: {
    paddingBottom: 10,
  },
  imageContainer: {
    borderRadius: 3,
    elevation: 1,
    overflow: "hidden",
  },
  image: {
    height: 122,
    width: "auto",
  },
  cardContent: {
    // play around with these values
    padding: 15,
    margin: 10,
  },
  card: {
    backgroundColor: "#FFFFFF",
    width: "100%",
    minWidth: "100%",
    marginBottom: 10,
    height: 120,
  },
  shadowProp: {
    shadowColor: "#171717",
    shadowOffset: { width: -2, height: 4 },
    shadowOpacity: 0.2,
    shadowRadius: 3,
  },
});

export default UpcomingEventsCard;

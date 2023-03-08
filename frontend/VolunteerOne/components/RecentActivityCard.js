import { StyleSheet } from "react-native";
import { Block, Text } from "galio-framework";

/*
Description:
  This component returns a recent activity card. Receives the data to insert into the card. 
Props received:
  data: date, description, likes, comments
*/

const RecentActivityCard = ({ data }) => {
  const { date, description, likes, comments } = data;
  return (
    <Block style={[styles.card, styles.shadowProp]}>
      <Text bold size={12} color="#525F7F" style={styles.textSpacing}>
        {date}
      </Text>
      <Text size={12} color="#525F7F" numberOfLines={3} style={styles.textSpacing}>
        {description}
      </Text>
      <Block row space="between">
        <Text size={12} color="#8898AA">
          {likes} likes
        </Text>
        <Text size={12} color="#8898AA">
          {comments} comments
        </Text>
      </Block>
    </Block>
  );
};

const styles = StyleSheet.create({
  textSpacing: {
    paddingBottom: 10
  },
  card: {
    backgroundColor: "#FFFFFF",
    width: "100%",
    minWidth: "100%",
    // borderRadius: 10,
    padding: 15,
    // margin: 10,
    marginBottom: 10,
    height: 120
  },
  shadowProp: {
    shadowColor: "#171717",
    shadowOffset: { width: -2, height: 4 },
    shadowOpacity: 0.2,
    shadowRadius: 3,
  },
});

export default RecentActivityCard;

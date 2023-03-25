// This page is temporarily being used to test other pages (notifications, friends, etc.) - matt
import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';
import { Button } from "../../components";
import NewAnouncementModal from "../../components/NewAnnouncementModal"
import NewPostModal from '../../components/NewPostModal';


// ================================= View Friends Page ================================= //

const { width } = Dimensions.get('screen');

class ViewFriendsPage extends React.Component {
  renderNotifications = () => {
    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.notifications}>
        
      <NewPostModal></NewPostModal>
       
      </ScrollView>
    )
  }

  render() {
    return (
      <Block flex center style={styles.home}>
        {this.renderNotifications()}
      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,    
  },
  notifications: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

export default ViewFriendsPage;
